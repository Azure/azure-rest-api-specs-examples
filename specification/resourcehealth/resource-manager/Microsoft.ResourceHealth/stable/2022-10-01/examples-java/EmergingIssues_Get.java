import com.azure.resourcemanager.resourcehealth.models.IssueNameParameter;

/** Samples for EmergingIssues Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2022-10-01/examples/EmergingIssues_Get.json
     */
    /**
     * Sample code: GetEmergingIssues.
     *
     * @param manager Entry point to ResourceHealthManager.
     */
    public static void getEmergingIssues(com.azure.resourcemanager.resourcehealth.ResourceHealthManager manager) {
        manager.emergingIssues().getWithResponse(IssueNameParameter.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
