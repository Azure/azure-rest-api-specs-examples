
import com.azure.resourcemanager.computelimit.models.SharedLimitCapProperties;

/**
 * Samples for SharedLimitCaps CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/SharedLimitCaps_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or update a shared limit cap for a VM family.
     * 
     * @param manager Entry point to ComputeLimitManager.
     */
    public static void
        createOrUpdateASharedLimitCapForAVMFamily(com.azure.resourcemanager.computelimit.ComputeLimitManager manager) {
        manager.sharedLimitCaps().define("StandardDSv3Family").withExistingLocation("eastus")
            .withProperties(new SharedLimitCapProperties().withDefaultMemberCap(100).withIsBoundedCap(true)).create();
    }
}
