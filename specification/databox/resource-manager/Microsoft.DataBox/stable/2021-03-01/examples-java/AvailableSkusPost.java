import com.azure.resourcemanager.databox.models.AvailableSkuRequest;
import com.azure.resourcemanager.databox.models.TransferType;

/** Samples for Service ListAvailableSkusByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2021-03-01/examples/AvailableSkusPost.json
     */
    /**
     * Sample code: AvailableSkusPost.
     *
     * @param manager Entry point to DataBoxManager.
     */
    public static void availableSkusPost(com.azure.resourcemanager.databox.DataBoxManager manager) {
        manager
            .services()
            .listAvailableSkusByResourceGroup(
                "bvttoolrg6",
                "westus",
                new AvailableSkuRequest()
                    .withTransferType(TransferType.IMPORT_TO_AZURE)
                    .withCountry("US")
                    .withLocation("westus"),
                com.azure.core.util.Context.NONE);
    }
}
