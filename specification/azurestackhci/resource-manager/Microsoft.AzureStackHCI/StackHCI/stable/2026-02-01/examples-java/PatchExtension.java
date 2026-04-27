
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.azurestackhci.models.Extension;
import com.azure.resourcemanager.azurestackhci.models.ExtensionPatchParameters;
import java.io.IOException;

/**
 * Samples for Extensions Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/PatchExtension.json
     */
    /**
     * Sample code: Update Arc Extension.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void updateArcExtension(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager)
        throws IOException {
        Extension resource = manager.extensions().getWithResponse("test-rg", "myCluster", "default",
            "MicrosoftMonitoringAgent", com.azure.core.util.Context.NONE).getValue();
        resource.update()
            .withExtensionParameters(
                new ExtensionPatchParameters().withTypeHandlerVersion("1.10").withEnableAutomaticUpgrade(false)
                    .withSettings(SerializerFactory.createDefaultManagementSerializerAdapter()
                        .deserialize("{\"workspaceId\":\"xx\"}", Object.class, SerializerEncoding.JSON))
                    .withProtectedSettings(SerializerFactory.createDefaultManagementSerializerAdapter()
                        .deserialize("{\"workspaceKey\":\"xx\"}", Object.class, SerializerEncoding.JSON)))
            .apply();
    }
}
