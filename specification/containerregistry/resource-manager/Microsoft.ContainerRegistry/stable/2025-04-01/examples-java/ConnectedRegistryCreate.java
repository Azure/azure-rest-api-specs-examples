
import com.azure.resourcemanager.containerregistry.fluent.models.ConnectedRegistryInner;
import com.azure.resourcemanager.containerregistry.models.ConnectedRegistryMode;
import com.azure.resourcemanager.containerregistry.models.GarbageCollectionProperties;
import com.azure.resourcemanager.containerregistry.models.ParentProperties;
import com.azure.resourcemanager.containerregistry.models.SyncProperties;
import java.time.Duration;
import java.util.Arrays;

/**
 * Samples for ConnectedRegistries Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2025-04-01/examples/
     * ConnectedRegistryCreate.json
     */
    /**
     * Sample code: ConnectedRegistryCreate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void connectedRegistryCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getConnectedRegistries().create("myResourceGroup",
            "myRegistry", "myConnectedRegistry",
            new ConnectedRegistryInner().withMode(ConnectedRegistryMode.READ_WRITE)
                .withParent(new ParentProperties().withSyncProperties(
                    new SyncProperties().withTokenId("fakeTokenPlaceholder").withSchedule("0 9 * * *")
                        .withSyncWindow(Duration.parse("PT3H")).withMessageTtl(Duration.parse("P2D"))))
                .withClientTokenIds(Arrays.asList(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client1Token"))
                .withNotificationsList(Arrays.asList("hello-world:*:*", "sample/repo/*:1.0:*"))
                .withGarbageCollection(new GarbageCollectionProperties().withEnabled(true).withSchedule("0 5 * * *")),
            com.azure.core.util.Context.NONE);
    }
}
