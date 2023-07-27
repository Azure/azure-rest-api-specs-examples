import com.azure.resourcemanager.batch.models.AccountKeyType;
import com.azure.resourcemanager.batch.models.BatchAccountRegenerateKeyParameters;

/** Samples for BatchAccount RegenerateKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/BatchAccountRegenerateKey.json
     */
    /**
     * Sample code: BatchAccountRegenerateKey.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void batchAccountRegenerateKey(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .batchAccounts()
            .regenerateKeyWithResponse(
                "default-azurebatch-japaneast",
                "sampleacct",
                new BatchAccountRegenerateKeyParameters().withKeyName(AccountKeyType.PRIMARY),
                com.azure.core.util.Context.NONE);
    }
}
