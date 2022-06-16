import com.azure.resourcemanager.synapse.models.BlobAuditingPolicyState;

/** Samples for SqlPoolBlobAuditingPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateSqlPoolBlobAuditingWithMinParameters.json
     */
    /**
     * Sample code: Create or update a database's blob auditing policy with minimal parameters.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void createOrUpdateADatabaseSBlobAuditingPolicyWithMinimalParameters(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolBlobAuditingPolicies()
            .define()
            .withExistingSqlPool("blobauditingtest-4799", "blobauditingtest-6440", "testdb")
            .withState(BlobAuditingPolicyState.ENABLED)
            .withStorageEndpoint("https://mystorage.blob.core.windows.net")
            .withStorageAccountAccessKey(
                "sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD==")
            .create();
    }
}
