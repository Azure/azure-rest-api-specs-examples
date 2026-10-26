
import com.azure.resourcemanager.datatransfer.models.Schema;

/**
 * Samples for AzureDataTransfer ValidateSchema.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/postValidateSchema.json
     */
    /**
     * Sample code: Performs action request.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void performsActionRequest(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.azureDataTransfers().validateSchemaWithResponse(new Schema().withName("schema.xml"),
            com.azure.core.util.Context.NONE);
    }
}
