/** Samples for Creators ListByAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/maps/resource-manager/Microsoft.Maps/stable/2023-06-01/examples/ListMapsCreatorsByAccount.json
     */
    /**
     * Sample code: List Creator Resources By Account.
     *
     * @param manager Entry point to AzureMapsManager.
     */
    public static void listCreatorResourcesByAccount(com.azure.resourcemanager.maps.AzureMapsManager manager) {
        manager.creators().listByAccount("myResourceGroup", "myMapsAccount", com.azure.core.util.Context.NONE);
    }
}
