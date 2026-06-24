
import com.azure.resourcemanager.computelimit.models.SetMemberCapOverridesRequest;
import java.util.Arrays;

/**
 * Samples for SharedLimitCaps SetMemberCapOverrides.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/SharedLimitCaps_SetMemberCapOverrides_ClearAll.json
     */
    /**
     * Sample code: Clear all member cap overrides (supply an empty array).
     * 
     * @param manager Entry point to ComputeLimitManager.
     */
    public static void clearAllMemberCapOverridesSupplyAnEmptyArray(
        com.azure.resourcemanager.computelimit.ComputeLimitManager manager) {
        manager.sharedLimitCaps().setMemberCapOverridesWithResponse("eastus", "StandardDSv3Family",
            new SetMemberCapOverridesRequest().withMemberCapOverrides(Arrays.asList()),
            com.azure.core.util.Context.NONE);
    }
}
