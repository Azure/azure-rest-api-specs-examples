```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.appplatform.fluent.models.BuildInner;
import com.azure.resourcemanager.appplatform.models.BuildProperties;
import java.util.HashMap;
import java.util.Map;

/** Samples for BuildService CreateOrUpdateBuild. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/BuildService_CreateOrUpdateBuild.json
     */
    /**
     * Sample code: BuildService_CreateOrUpdateBuild.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void buildServiceCreateOrUpdateBuild(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getBuildServices()
            .createOrUpdateBuildWithResponse(
                "myResourceGroup",
                "myservice",
                "default",
                "mybuild",
                new BuildInner()
                    .withProperties(
                        new BuildProperties()
                            .withRelativePath(
                                "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855-20210601-3ed9f4a2-986b-4bbd-b833-a42dccb2f777")
                            .withBuilder(
                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/buildServices/default/builders/default")
                            .withAgentPool(
                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/buildServices/default/agentPools/default")
                            .withEnv(mapOf("environmentVariable", "test"))),
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
