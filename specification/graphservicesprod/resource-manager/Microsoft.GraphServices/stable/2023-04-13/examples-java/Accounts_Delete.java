/** Samples for Accounts Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/graphservicesprod/resource-manager/Microsoft.GraphServices/stable/2023-04-13/examples/Accounts_Delete.json
     */
    /**
     * Sample code: Delete account resource.
     *
     * @param manager Entry point to GraphServicesManager.
     */
    public static void deleteAccountResource(com.azure.resourcemanager.graphservices.GraphServicesManager manager) {
        manager
            .accounts()
            .deleteByResourceGroupWithResponse(
                "testResourceGroupGRAM", "11111111-aaaa-1111-bbbb-111111111111", com.azure.core.util.Context.NONE);
    }
}
