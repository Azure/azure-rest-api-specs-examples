
import com.azure.resourcemanager.keyvault.fluent.models.ManagedHsmKeyProperties;
import com.azure.resourcemanager.keyvault.models.JsonWebKeyType;
import com.azure.resourcemanager.keyvault.models.ManagedHsmKeyCreateParameters;

/**
 * Samples for ManagedHsmKeys CreateIfNotExist.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/managedHsmCreateKey.json
     */
    /**
     * Sample code: Create a key.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getManagedHsmKeys().createIfNotExistWithResponse(
            "sample-group", "sample-managedhsm-name", "sample-key-name", new ManagedHsmKeyCreateParameters()
                .withProperties(new ManagedHsmKeyProperties().withKty(JsonWebKeyType.RSA)),
            com.azure.core.util.Context.NONE);
    }
}
