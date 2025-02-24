
import com.azure.resourcemanager.netapp.models.EncryptionTransitionRequest;

/**
 * Samples for Accounts TransitionToCmk.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-09-01/examples/
     * Accounts_TransitionEncryptionKey.json
     */
    /**
     * Sample code: Accounts_MigrateEncryptionKey.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void accountsMigrateEncryptionKey(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.accounts().transitionToCmk("myRG", "account1", new EncryptionTransitionRequest().withVirtualNetworkId(
            "/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/resourceGroups/myRG/providers/Microsoft.Network/virtualNetworks/vnet1")
            .withPrivateEndpointId(
                "/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/resourceGroups/myRG/providers/Microsoft.Network/privateEndpoints/privip1"),
            com.azure.core.util.Context.NONE);
    }
}
