
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.datafactory.models.AzureStorageLinkedService;
import com.azure.resourcemanager.datafactory.models.LinkedServiceResource;
import java.io.IOException;

/**
 * Samples for LinkedServices CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/LinkedServices_Update
     * .json
     */
    /**
     * Sample code: LinkedServices_Update.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void linkedServicesUpdate(com.azure.resourcemanager.datafactory.DataFactoryManager manager)
        throws IOException {
        LinkedServiceResource resource = manager.linkedServices().getWithResponse("exampleResourceGroup",
            "exampleFactoryName", "exampleLinkedService", null, com.azure.core.util.Context.NONE).getValue();
        resource.update().withProperties(new AzureStorageLinkedService().withDescription("Example description")
            .withConnectionString(SerializerFactory.createDefaultManagementSerializerAdapter().deserialize(
                "{\"type\":\"SecureString\",\"value\":\"DefaultEndpointsProtocol=https;AccountName=examplestorageaccount;AccountKey=<storage key>\"}",
                Object.class, SerializerEncoding.JSON)))
            .apply();
    }
}
