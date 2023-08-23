/** Samples for Accounts ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/maps/resource-manager/Microsoft.Maps/stable/2023-06-01/examples/ListAccountsByResourceGroup.json
     */
    /**
     * Sample code: List Accounts By Resource Group.
     *
     * @param manager Entry point to AzureMapsManager.
     */
    public static void listAccountsByResourceGroup(com.azure.resourcemanager.maps.AzureMapsManager manager) {
        manager.accounts().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
