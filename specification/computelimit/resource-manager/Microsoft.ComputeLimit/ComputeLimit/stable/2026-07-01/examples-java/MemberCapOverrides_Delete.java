
/**
 * Samples for MemberCapOverrides Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/MemberCapOverrides_Delete.json
     */
    /**
     * Sample code: Delete a single member cap override.
     * 
     * @param manager Entry point to ComputeLimitManager.
     */
    public static void
        deleteASingleMemberCapOverride(com.azure.resourcemanager.computelimit.ComputeLimitManager manager) {
        manager.memberCapOverrides().deleteWithResponse("eastus", "StandardDSv3Family",
            "11111111-1111-1111-1111-111111111111", com.azure.core.util.Context.NONE);
    }
}
