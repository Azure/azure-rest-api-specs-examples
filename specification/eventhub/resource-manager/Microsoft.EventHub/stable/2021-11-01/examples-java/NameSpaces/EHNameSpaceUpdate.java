import com.azure.core.util.Context;
import com.azure.resourcemanager.eventhubs.fluent.models.EHNamespaceInner;
import com.azure.resourcemanager.eventhubs.models.Identity;
import com.azure.resourcemanager.eventhubs.models.ManagedServiceIdentityType;
import java.util.HashMap;
import java.util.Map;

/** Samples for Namespaces Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/NameSpaces/EHNameSpaceUpdate.json
     */
    /**
     * Sample code: NamespacesUpdate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void namespacesUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .eventHubs()
            .manager()
            .serviceClient()
            .getNamespaces()
            .updateWithResponse(
                "ResurceGroupSample",
                "NamespaceSample",
                new EHNamespaceInner()
                    .withLocation("East US")
                    .withIdentity(
                        new Identity()
                            .withType(ManagedServiceIdentityType.SYSTEM_ASSIGNED_USER_ASSIGNED)
                            .withUserAssignedIdentities(
                                mapOf(
                                    "/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ud2",
                                    null))),
                Context.NONE);
    }

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
