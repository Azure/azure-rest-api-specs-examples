import com.azure.core.util.Context;
import com.azure.resourcemanager.appplatform.fluent.models.BindingResourceInner;
import com.azure.resourcemanager.appplatform.models.BindingResourceProperties;
import java.util.HashMap;
import java.util.Map;

/** Samples for Bindings CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/Bindings_CreateOrUpdate.json
     */
    /**
     * Sample code: Bindings_CreateOrUpdate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void bindingsCreateOrUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getBindings()
            .createOrUpdate(
                "myResourceGroup",
                "myservice",
                "myapp",
                "mybinding",
                new BindingResourceInner()
                    .withProperties(
                        new BindingResourceProperties()
                            .withResourceId(
                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DocumentDB/databaseAccounts/my-cosmosdb-1")
                            .withKey("xxxx")
                            .withBindingParameters(mapOf("apiType", "SQL", "databaseName", "db1"))),
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
