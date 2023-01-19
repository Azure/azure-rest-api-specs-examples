/** Samples for Accounts Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/maps/resource-manager/Microsoft.Maps/stable/2021-02-01/examples/DeleteAccount.json
     */
    /**
     * Sample code: DeleteAccount.
     *
     * @param manager Entry point to AzureMapsManager.
     */
    public static void deleteAccount(com.azure.resourcemanager.maps.AzureMapsManager manager) {
        manager
            .accounts()
            .deleteByResourceGroupWithResponse("myResourceGroup", "myMapsAccount", com.azure.core.util.Context.NONE);
    }
}
