
import com.azure.resourcemanager.netapp.models.EncryptionTransitionRequest;

/**
 * Samples for Accounts TransitionToCmk.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-01-01/examples/
     * Accounts_TransitionEncryptionKey.json
     */
    /**
     * Sample code: Accounts_MigrateEncryptionKey.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void accountsMigrateEncryptionKey(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.accounts().transitionToCmk("myRG", "account1", new EncryptionTransitionRequest().withVirtualNetworkId(
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.Network/virtualNetworks/vnet1")
            .withPrivateEndpointId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.Network/privateEndpoints/privip1"),
            com.azure.core.util.Context.NONE);
    }
}
