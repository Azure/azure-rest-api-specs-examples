
import com.azure.resourcemanager.storage.fluent.models.TableInner;
import com.azure.resourcemanager.storage.models.TableAccessPolicy;
import com.azure.resourcemanager.storage.models.TableSignedIdentifier;
import java.time.OffsetDateTime;
import java.util.Arrays;

/**
 * Samples for Table Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/TableOperationPutOrPatchAcls.
     * json
     */
    /**
     * Sample code: TableOperationPutOrPatchAcls.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void tableOperationPutOrPatchAcls(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getTables()
            .createWithResponse("res3376", "sto328", "table6185", new TableInner().withSignedIdentifiers(Arrays.asList(
                new TableSignedIdentifier().withId("MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI")
                    .withAccessPolicy(new TableAccessPolicy()
                        .withStartTime(OffsetDateTime.parse("2022-03-17T08:49:37.0000000Z"))
                        .withExpiryTime(OffsetDateTime.parse("2022-03-20T08:49:37.0000000Z")).withPermission("raud")),
                new TableSignedIdentifier().withId("PTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODklMTI")
                    .withAccessPolicy(new TableAccessPolicy()
                        .withStartTime(OffsetDateTime.parse("2022-03-17T08:49:37.0000000Z"))
                        .withExpiryTime(OffsetDateTime.parse("2022-03-20T08:49:37.0000000Z")).withPermission("rad")))),
                com.azure.core.util.Context.NONE);
    }
}
