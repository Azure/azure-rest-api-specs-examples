/** Samples for EmergingIssues List. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2022-10-01/examples/EmergingIssues_List.json
     */
    /**
     * Sample code: GetEmergingIssues.
     *
     * @param manager Entry point to ResourceHealthManager.
     */
    public static void getEmergingIssues(com.azure.resourcemanager.resourcehealth.ResourceHealthManager manager) {
        manager.emergingIssues().list(com.azure.core.util.Context.NONE);
    }
}
