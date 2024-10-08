
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import java.io.IOException;

/**
 * Samples for Extensions Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/
     * PutExtension.json
     */
    /**
     * Sample code: Create Arc Extension.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void createArcExtension(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager)
        throws IOException {
        manager.extensions().define("MicrosoftMonitoringAgent")
            .withExistingArcSetting("test-rg", "myCluster", "default").withPublisher("Microsoft.Compute")
            .withTypePropertiesType("MicrosoftMonitoringAgent").withTypeHandlerVersion("1.10")
            .withSettings(SerializerFactory.createDefaultManagementSerializerAdapter()
                .deserialize("{\"workspaceId\":\"xx\"}", Object.class, SerializerEncoding.JSON))
            .withProtectedSettings(SerializerFactory.createDefaultManagementSerializerAdapter()
                .deserialize("{\"workspaceKey\":\"xx\"}", Object.class, SerializerEncoding.JSON))
            .withEnableAutomaticUpgrade(false).create();
    }
}
