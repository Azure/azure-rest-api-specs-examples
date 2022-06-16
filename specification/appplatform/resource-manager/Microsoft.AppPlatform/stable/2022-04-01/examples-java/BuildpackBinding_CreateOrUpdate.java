import com.azure.core.util.Context;
import com.azure.resourcemanager.appplatform.fluent.models.BuildpackBindingResourceInner;
import com.azure.resourcemanager.appplatform.models.BindingType;
import com.azure.resourcemanager.appplatform.models.BuildpackBindingLaunchProperties;
import com.azure.resourcemanager.appplatform.models.BuildpackBindingProperties;
import java.util.HashMap;
import java.util.Map;

/** Samples for BuildpackBinding CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/BuildpackBinding_CreateOrUpdate.json
     */
    /**
     * Sample code: BuildpackBinding_CreateOrUpdate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void buildpackBindingCreateOrUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getBuildpackBindings()
            .createOrUpdate(
                "myResourceGroup",
                "myservice",
                "default",
                "default",
                "myBuildpackBinding",
                new BuildpackBindingResourceInner()
                    .withProperties(
                        new BuildpackBindingProperties()
                            .withBindingType(BindingType.APPLICATION_INSIGHTS)
                            .withLaunchProperties(
                                new BuildpackBindingLaunchProperties()
                                    .withProperties(
                                        mapOf("abc", "def", "any-string", "any-string", "sampling-rate", "12.0"))
                                    .withSecrets(
                                        mapOf(
                                            "connection-string",
                                            "XXXXXXXXXXXXXXXXX=XXXXXXXXXXXXX-XXXXXXXXXXXXXXXXXXX;XXXXXXXXXXXXXXXXX=XXXXXXXXXXXXXXXXXXX")))),
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
