
import com.azure.resourcemanager.containerregistry.models.AuditLogStatus;
import com.azure.resourcemanager.containerregistry.models.ConnectedRegistryUpdateParameters;
import com.azure.resourcemanager.containerregistry.models.GarbageCollectionProperties;
import com.azure.resourcemanager.containerregistry.models.LogLevel;
import com.azure.resourcemanager.containerregistry.models.LoggingProperties;
import com.azure.resourcemanager.containerregistry.models.SyncUpdateProperties;
import java.time.Duration;
import java.util.Arrays;

/**
 * Samples for ConnectedRegistries Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2025-04-01/examples/
     * ConnectedRegistryUpdate.json
     */
    /**
     * Sample code: ConnectedRegistryUpdate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void connectedRegistryUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getConnectedRegistries()
            .update("myResourceGroup", "myRegistry", "myScopeMap", new ConnectedRegistryUpdateParameters()
                .withSyncProperties(new SyncUpdateProperties().withSchedule("0 0 */10 * *")
                    .withSyncWindow(Duration.parse("P2D")).withMessageTtl(Duration.parse("P30D")))
                .withLogging(
                    new LoggingProperties().withLogLevel(LogLevel.DEBUG).withAuditLogStatus(AuditLogStatus.ENABLED))
                .withClientTokenIds(Arrays.asList(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client1Token",
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client2Token"))
                .withNotificationsList(Arrays.asList("hello-world:*:*", "sample/repo/*:1.0:*"))
                .withGarbageCollection(new GarbageCollectionProperties().withEnabled(true).withSchedule("0 5 * * *")),
                com.azure.core.util.Context.NONE);
    }
}
