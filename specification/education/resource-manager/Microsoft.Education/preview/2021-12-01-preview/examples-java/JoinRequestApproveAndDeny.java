import com.azure.core.util.Context;

/** Samples for JoinRequests Approve. */
public final class Main {
    /*
     * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/JoinRequestApproveAndDeny.json
     */
    /**
     * Sample code: JoinRequestApprove.
     *
     * @param manager Entry point to EducationManager.
     */
    public static void joinRequestApprove(com.azure.resourcemanager.education.EducationManager manager) {
        manager
            .joinRequests()
            .approveWithResponse(
                "{billingAccountName}",
                "{billingProfileName}",
                "{invoiceSectionName}",
                "{joinRequestName}",
                Context.NONE);
    }
}
