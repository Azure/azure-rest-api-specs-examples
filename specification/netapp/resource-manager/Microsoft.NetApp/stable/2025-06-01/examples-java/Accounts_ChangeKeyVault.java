
import com.azure.resourcemanager.netapp.models.ChangeKeyVault;
import com.azure.resourcemanager.netapp.models.KeyVaultPrivateEndpoint;
import java.util.Arrays;

/**
 * Samples for Accounts ChangeKeyVault.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-06-01/examples/Accounts_ChangeKeyVault.json
     */
    /**
     * Sample code: Accounts_ChangeKeyVault.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void accountsChangeKeyVault(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.accounts().changeKeyVault("myRG", "account1", new ChangeKeyVault()
            .withKeyVaultUri("fakeTokenPlaceholder").withKeyName("fakeTokenPlaceholder")
            .withKeyVaultResourceId("fakeTokenPlaceholder")
            .withKeyVaultPrivateEndpoints(Arrays.asList(new KeyVaultPrivateEndpoint().withVirtualNetworkId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.Network/virtualNetworks/vnet1")
                .withPrivateEndpointId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.Network/privateEndpoints/privip1"))),
            com.azure.core.util.Context.NONE);
    }
}
