import org.jsoup.Jsoup;
import org.jsoup.nodes.Document;
import org.jsoup.nodes.Element;
import org.jsoup.select.Elements;

import java.util.ArrayList;

public class Crawler {
    public static void main(String[] args) throws Exception {
        String url = "https://commoncrawl.org/";

        // 1. Fetch + parse HTML into DOM
        // 2. Select all <a> tags with href attributes
        Document doc = Jsoup.connect(url).get();

        
        Elements links = doc.select("a[href]");

        ArrayList<String> linkList = new ArrayList<>();

        for (Element link : links) {
            String absUrl = link.absUrl("href"); // handles relative → absolute
            linkList.add(absUrl);
        }

        for (String l : linkList) {
            System.out.println(l);
        }
    }
}