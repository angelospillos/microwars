package naughty.dtos;

public class StatusResponse {
    private String status = "ok";

    public StatusResponse() {
    }

    public StatusResponse(final String status) {
        this.status = status;
    }

    public String getStatus() {
        return status;
    }

    public void setStatus(final String status) {
        this.status = status;
    }
}
