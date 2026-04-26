
import java.awt.Desktop;
import java.io.*;
import java.net.URI;

public class WebPager {

    public static void main(String[] args) {
        Desktop desk = Desktop.getDesktop();
      
        // now we enter our URL that we want to open in our
        // default browser
        try {
            desk.browse(new URI("https://bloomberg.com/"));
        }
        catch(Exception e){
            System.out.println("error in opening web page");
        }



    }
}