import com.azure.core.util.Context;

/** Samples for JoinRequests Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/JoinRequest.json
     */
    /**
     * Sample code: JoinRequest.
     *
     * @param manager Entry point to EducationManager.
     */
    public static void joinRequest(com.azure.resourcemanager.education.EducationManager manager) {
        manager
            .joinRequests()
            .getWithResponse(
                "{billingAccountName}",
                "{billingProfileName}",
                "{invoiceSectionName}",
                "{joinRequestName}",
                Context.NONE);
    }
}
