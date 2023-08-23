/** Samples for Accounts GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/maps/resource-manager/Microsoft.Maps/stable/2023-06-01/examples/GetAccount.json
     */
    /**
     * Sample code: GetAccount.
     *
     * @param manager Entry point to AzureMapsManager.
     */
    public static void getAccount(com.azure.resourcemanager.maps.AzureMapsManager manager) {
        manager
            .accounts()
            .getByResourceGroupWithResponse("myResourceGroup", "myMapsAccount", com.azure.core.util.Context.NONE);
    }
}
