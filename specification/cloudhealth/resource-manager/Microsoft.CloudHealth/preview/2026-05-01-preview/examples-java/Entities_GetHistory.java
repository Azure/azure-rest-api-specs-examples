
import com.azure.resourcemanager.cloudhealth.models.EntityHistoryRequest;
import java.time.OffsetDateTime;

/**
 * Samples for Entities GetHistory.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Entities_GetHistory.json
     */
    /**
     * Sample code: Entities_GetHistory.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void entitiesGetHistory(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.entities().getHistoryWithResponse("online-store-rg", "online-store", "web-frontend",
            new EntityHistoryRequest().withStartAt(OffsetDateTime.parse("2026-05-03T09:30:00Z"))
                .withEndAt(OffsetDateTime.parse("2026-05-04T09:30:00Z")).withTop(100),
            com.azure.core.util.Context.NONE);
    }
}
