
import com.azure.resourcemanager.cloudhealth.models.SignalHistoryRequest;
import java.time.OffsetDateTime;

/**
 * Samples for Entities GetSignalHistory.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/Entities_GetSignalHistory.json
     */
    /**
     * Sample code: Entities_GetSignalHistory.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void entitiesGetSignalHistory(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.entities().getSignalHistoryWithResponse("rgopenapi", "myHealthModel", "entity1",
            new SignalHistoryRequest().withSignalName("uniqueSignalName1")
                .withStartAt(OffsetDateTime.parse("2025-12-11T10:00:00Z"))
                .withEndAt(OffsetDateTime.parse("2025-12-12T10:00:00Z")),
            com.azure.core.util.Context.NONE);
    }
}
