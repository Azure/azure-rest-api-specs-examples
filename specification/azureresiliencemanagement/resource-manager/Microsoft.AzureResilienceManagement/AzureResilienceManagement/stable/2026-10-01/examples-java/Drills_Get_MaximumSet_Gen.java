
/**
 * Samples for Drills Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Drills_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Drills_Get_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void
        drillsGetMaximumSet(com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.drills().getWithResponse("sampleServiceGroupName", "drill1", com.azure.core.util.Context.NONE);
    }
}
