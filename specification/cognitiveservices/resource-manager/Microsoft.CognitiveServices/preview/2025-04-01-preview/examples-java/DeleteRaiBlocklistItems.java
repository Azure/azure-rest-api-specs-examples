
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import java.io.IOException;

/**
 * Samples for RaiBlocklistItems BatchDelete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2025-04-01-preview/examples/
     * DeleteRaiBlocklistItems.json
     */
    /**
     * Sample code: DeleteRaiBlocklistItems.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void deleteRaiBlocklistItems(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) throws IOException {
        manager.raiBlocklistItems().batchDeleteWithResponse("resourceGroupName", "accountName", "raiBlocklistName",
            SerializerFactory.createDefaultManagementSerializerAdapter()
                .deserialize("[\"myblocklistitem1\",\"myblocklistitem2\"]", Object.class, SerializerEncoding.JSON),
            com.azure.core.util.Context.NONE);
    }
}
