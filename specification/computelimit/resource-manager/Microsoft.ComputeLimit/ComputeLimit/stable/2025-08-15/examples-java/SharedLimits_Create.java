
import com.azure.resourcemanager.computelimit.models.SharedLimitProperties;

/**
 * Samples for SharedLimits Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-15/SharedLimits_Create.json
     */
    /**
     * Sample code: Create a shared limit.
     * 
     * @param manager Entry point to ComputeLimitManager.
     */
    public static void createASharedLimit(com.azure.resourcemanager.computelimit.ComputeLimitManager manager) {
        manager.sharedLimits().define("StandardDSv3Family").withExistingLocation("eastus")
            .withProperties(new SharedLimitProperties()).create();
    }
}
