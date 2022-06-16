import com.azure.core.util.Context;
import com.azure.resourcemanager.batch.models.AutoStorageBaseProperties;
import com.azure.resourcemanager.batch.models.BatchAccount;

/** Samples for BatchAccount Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/BatchAccountUpdate.json
     */
    /**
     * Sample code: BatchAccountUpdate.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void batchAccountUpdate(com.azure.resourcemanager.batch.BatchManager manager) {
        BatchAccount resource =
            manager
                .batchAccounts()
                .getByResourceGroupWithResponse("default-azurebatch-japaneast", "sampleacct", Context.NONE)
                .getValue();
        resource
            .update()
            .withAutoStorage(
                new AutoStorageBaseProperties()
                    .withStorageAccountId(
                        "/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Storage/storageAccounts/samplestorage"))
            .apply();
    }
}
