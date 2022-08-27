import com.azure.core.util.Context;
import com.azure.resourcemanager.storage.fluent.models.LegalHoldInner;
import java.util.Arrays;

/** Samples for BlobContainers SetLegalHold. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/BlobContainersSetLegalHoldAllowProtectedAppendWritesAll.json
     */
    /**
     * Sample code: SetLegalHoldContainersWithAllowProtectedAppendWritesAll.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void setLegalHoldContainersWithAllowProtectedAppendWritesAll(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getBlobContainers()
            .setLegalHoldWithResponse(
                "res4303",
                "sto7280",
                "container8723",
                new LegalHoldInner()
                    .withTags(Arrays.asList("tag1", "tag2", "tag3"))
                    .withAllowProtectedAppendWritesAll(true),
                Context.NONE);
    }
}
