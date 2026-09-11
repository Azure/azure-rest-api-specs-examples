
import com.azure.resourcemanager.containerregistry.models.AuditLogStatus;
import com.azure.resourcemanager.containerregistry.models.ConnectedRegistryUpdateParameters;
import com.azure.resourcemanager.containerregistry.models.GarbageCollectionProperties;
import com.azure.resourcemanager.containerregistry.models.LogLevel;
import com.azure.resourcemanager.containerregistry.models.LoggingProperties;
import com.azure.resourcemanager.containerregistry.models.SyncUpdateProperties;
import java.time.Duration;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ConnectedRegistries Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01-preview/ConnectedRegistryUpdate.json
     */
    /**
     * Sample code: ConnectedRegistryUpdate.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void
        connectedRegistryUpdate(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getConnectedRegistries().update("myResourceGroup", "myRegistry", "myScopeMap",
            new ConnectedRegistryUpdateParameters()
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

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
