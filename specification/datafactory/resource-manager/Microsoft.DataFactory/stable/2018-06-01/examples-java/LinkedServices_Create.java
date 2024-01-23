
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.datafactory.models.AzureStorageLinkedService;
import java.io.IOException;

/**
 * Samples for LinkedServices CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/LinkedServices_Create
     * .json
     */
    /**
     * Sample code: LinkedServices_Create.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void linkedServicesCreate(com.azure.resourcemanager.datafactory.DataFactoryManager manager)
        throws IOException {
        manager.linkedServices().define("exampleLinkedService")
            .withExistingFactory("exampleResourceGroup", "exampleFactoryName")
            .withProperties(new AzureStorageLinkedService()
                .withConnectionString(SerializerFactory.createDefaultManagementSerializerAdapter().deserialize(
                    "{\"type\":\"SecureString\",\"value\":\"DefaultEndpointsProtocol=https;AccountName=examplestorageaccount;AccountKey=<storage key>\"}",
                    Object.class, SerializerEncoding.JSON)))
            .create();
    }
}
