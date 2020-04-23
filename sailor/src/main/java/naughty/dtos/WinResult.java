package naughty.dtos;

import io.micronaut.core.annotation.Introspected;

import java.util.Date;

@Introspected
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
