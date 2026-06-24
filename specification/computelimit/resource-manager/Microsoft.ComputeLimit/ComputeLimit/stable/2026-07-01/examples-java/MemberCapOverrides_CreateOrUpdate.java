
import com.azure.resourcemanager.computelimit.models.MemberCapOverrideProperties;

/**
 * Samples for MemberCapOverrides CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/MemberCapOverrides_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or update a single member cap override.
     * 
     * @param manager Entry point to ComputeLimitManager.
     */
    public static void
        createOrUpdateASingleMemberCapOverride(com.azure.resourcemanager.computelimit.ComputeLimitManager manager) {
        manager.memberCapOverrides().define("11111111-1111-1111-1111-111111111111")
            .withExistingSharedLimitCap("eastus", "StandardDSv3Family")
            .withProperties(new MemberCapOverrideProperties().withCap(250)).create();
    }
}
