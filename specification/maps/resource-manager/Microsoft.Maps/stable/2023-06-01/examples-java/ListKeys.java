/** Samples for Accounts ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/maps/resource-manager/Microsoft.Maps/stable/2023-06-01/examples/ListKeys.json
     */
    /**
     * Sample code: List Keys.
     *
     * @param manager Entry point to AzureMapsManager.
     */
    public static void listKeys(com.azure.resourcemanager.maps.AzureMapsManager manager) {
        manager.accounts().listKeysWithResponse("myResourceGroup", "myMapsAccount", com.azure.core.util.Context.NONE);
    }
}
