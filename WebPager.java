
import java.awt.Desktop;
import java.io.*;
import java.net.URI;
import java.net.URL;
import java.net.URLConnection;
import java.util.ArrayList;
import java.util.Scanner;
import java.util.regex.Matcher;
import java.util.regex.Pattern;
import java.nio.file.Files;
import java.nio.file.Paths;

public class WebPager {

    public static void main(String[] args) throws FileNotFoundException, IOException {
        // Desktop desk = Desktop.getDesktop();
      
        // now we enter our URL that we want to open in our
        // default browser
        // try {
        //     desk.browse(new URI("https://bloomberg.com/"));
        // }
        // catch(Exception e){
        //     System.out.println("error in opening web page");
        // }


        String content = null;
        URLConnection connection = null;
        try {
        
            connection =  new URL("https://www.fluentwithstories.com/stories/es").openConnection();
            Scanner scanner = new Scanner(connection.getInputStream());
            scanner.useDelimiter("\\Z");
            content = scanner.next();
            scanner.close();
            Files.writeString(Paths.get("output.txt"), content);

        }catch ( Exception ex ) {
            ex.printStackTrace();
        }

        
        // try (BufferedReader reader = new BufferedReader(new FileReader("output.txt"))) {
        //     String line;


        //     while ((line = reader.readLine()) != null) {
                
        //     }
        // }

        // Somethings in the water !! Need to look into this again
        String HTMLPage = "output.txt";
        Pattern linkPattern = Pattern.compile("(<a[^>]+>.+?<a>)",  Pattern.CASE_INSENSITIVE|Pattern.DOTALL);
        Matcher pageMatcher = linkPattern.matcher(HTMLPage);
        ArrayList<String> links = new ArrayList<String>();
        while(pageMatcher.find()){
            links.add(pageMatcher.group());
        }

        for (String link : links) {
            System.out.println(link);
        }

    }
}