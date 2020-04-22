package naughty.dtos;

import java.util.Date;

public class WinResult {
    private final String name = "Naughty Sailor";
    private final Date date = new Date();

    public String getName() {
        return name;
    }

    public Date getDate() {
        return date;
    }
}
