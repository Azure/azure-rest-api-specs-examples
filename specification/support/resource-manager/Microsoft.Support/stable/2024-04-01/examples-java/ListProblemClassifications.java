
/**
 * Samples for ProblemClassifications List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/ListProblemClassifications.
     * json
     */
    /**
     * Sample code: Gets list of problemClassifications for a service for which a support ticket can be created.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void getsListOfProblemClassificationsForAServiceForWhichASupportTicketCanBeCreated(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager.problemClassifications().list("service_guid", com.azure.core.util.Context.NONE);
    }
}
