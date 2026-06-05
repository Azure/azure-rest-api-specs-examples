
import com.azure.resourcemanager.trafficmanager.fluent.models.ProfileInner;
import com.azure.resourcemanager.trafficmanager.models.RecordType;

/**
 * Samples for Profiles Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-04-01-preview/Profile-PATCH-RecordType.json
     */
    /**
     * Sample code: Profile-PATCH-RecordType.
     * 
     * @param manager Entry point to TrafficManager.
     */
    public static void profilePATCHRecordType(com.azure.resourcemanager.trafficmanager.TrafficManager manager) {
        manager.serviceClient().getProfiles().updateWithResponse("azuresdkfornetautoresttrafficmanager2583",
            "azuresdkfornetautoresttrafficmanager6192", new ProfileInner().withRecordType(RecordType.CNAME),
            com.azure.core.util.Context.NONE);
    }
}
