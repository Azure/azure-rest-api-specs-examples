
import com.azure.resourcemanager.cosmos.models.ClientEncryptionKeyCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.ClientEncryptionKeyResource;
import com.azure.resourcemanager.cosmos.models.KeyWrapMetadata;

/**
 * Samples for SqlResources CreateUpdateClientEncryptionKey.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBSqlClientEncryptionKeyCreateUpdate.json
     */
    /**
     * Sample code: CosmosDBClientEncryptionKeyCreateUpdate.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBClientEncryptionKeyCreateUpdate(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources()
            .createUpdateClientEncryptionKey("rgName", "accountName", "databaseName", "cekName",
                new ClientEncryptionKeyCreateUpdateParameters().withResource(new ClientEncryptionKeyResource()
                    .withId("cekName").withEncryptionAlgorithm("AEAD_AES_256_CBC_HMAC_SHA256")
                    .withWrappedDataEncryptionKey("U3dhZ2dlciByb2Nrcw==".getBytes())
                    .withKeyWrapMetadata(new KeyWrapMetadata().withName("customerManagedKey").withType("AzureKeyVault")
                        .withValue("AzureKeyVault Key URL").withAlgorithm("RSA-OAEP"))),
                com.azure.core.util.Context.NONE);
    }
}
