/** Samples for Services List. */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/ListServices.json
     */
    /**
     * Sample code: Gets list of services for which a support ticket can be created.
     *
     * @param manager Entry point to SupportManager.
     */
    public static void getsListOfServicesForWhichASupportTicketCanBeCreated(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager.services().list(com.azure.core.util.Context.NONE);
    }
}
