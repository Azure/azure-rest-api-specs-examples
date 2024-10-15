
/**
 * Samples for JoinRequests Deny.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/
     * JoinRequestApproveAndDeny.json
     */
    /**
     * Sample code: JoinRequestDeny.
     * 
     * @param manager Entry point to EducationManager.
     */
    public static void joinRequestDeny(com.azure.resourcemanager.education.EducationManager manager) {
        manager.joinRequests().denyWithResponse("{billingAccountName}", "{billingProfileName}", "{invoiceSectionName}",
            "{joinRequestName}", com.azure.core.util.Context.NONE);
    }
}
