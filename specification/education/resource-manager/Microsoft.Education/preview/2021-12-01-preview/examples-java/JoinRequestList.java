import com.azure.core.util.Context;

/** Samples for JoinRequests List. */
public final class Main {
    /*
     * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/JoinRequestList.json
     */
    /**
     * Sample code: JoinRequestList.
     *
     * @param manager Entry point to EducationManager.
     */
    public static void joinRequestList(com.azure.resourcemanager.education.EducationManager manager) {
        manager
            .joinRequests()
            .list("{billingAccountName}", "{billingProfileName}", "{invoiceSectionName}", false, Context.NONE);
    }
}
