
import com.azure.resourcemanager.keyvault.fluent.models.ManagedHsmKeyProperties;
import com.azure.resourcemanager.keyvault.models.JsonWebKeyType;
import com.azure.resourcemanager.keyvault.models.ManagedHsmKeyCreateParameters;

/**
 * Samples for ManagedHsmKeys CreateIfNotExist.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/managedHsmCreateKey.json
     */
    /**
     * Sample code: Create a key.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void createAKey(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getManagedHsmKeys().createIfNotExistWithResponse(
            "sample-group", "sample-managedhsm-name", "sample-key-name", new ManagedHsmKeyCreateParameters()
                .withProperties(new ManagedHsmKeyProperties().withKty(JsonWebKeyType.RSA)),
            com.azure.core.util.Context.NONE);
    }
}
