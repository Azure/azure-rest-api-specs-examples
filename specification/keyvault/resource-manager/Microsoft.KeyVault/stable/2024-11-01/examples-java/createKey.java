
import com.azure.resourcemanager.keyvault.fluent.models.KeyProperties;
import com.azure.resourcemanager.keyvault.models.JsonWebKeyType;
import com.azure.resourcemanager.keyvault.models.KeyCreateParameters;

/**
 * Samples for Keys CreateIfNotExist.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2024-11-01/examples/createKey.json
     */
    /**
     * Sample code: Create a key.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getKeys().createIfNotExistWithResponse("sample-group",
            "sample-vault-name", "sample-key-name",
            new KeyCreateParameters().withProperties(new KeyProperties().withKty(JsonWebKeyType.RSA)),
            com.azure.core.util.Context.NONE);
    }
}
