import com.azure.resourcemanager.appcontainers.fluent.models.DaprComponentInner;
import com.azure.resourcemanager.appcontainers.models.DaprMetadata;
import com.azure.resourcemanager.appcontainers.models.Secret;
import java.util.Arrays;

/** Samples for DaprComponents CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-11-01-preview/examples/DaprComponents_CreateOrUpdate_Secrets.json
     */
    /**
     * Sample code: Create or update dapr component with secrets.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void createOrUpdateDaprComponentWithSecrets(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .daprComponents()
            .createOrUpdateWithResponse(
                "examplerg",
                "myenvironment",
                "reddog",
                new DaprComponentInner()
                    .withComponentType("state.azure.cosmosdb")
                    .withVersion("v1")
                    .withIgnoreErrors(false)
                    .withInitTimeout("50s")
                    .withSecrets(Arrays.asList(new Secret().withName("masterkey").withValue("keyvalue")))
                    .withMetadata(
                        Arrays
                            .asList(
                                new DaprMetadata().withName("url").withValue("<COSMOS-URL>"),
                                new DaprMetadata().withName("database").withValue("itemsDB"),
                                new DaprMetadata().withName("collection").withValue("items"),
                                new DaprMetadata().withName("masterkey").withSecretRef("fakeTokenPlaceholder")))
                    .withScopes(Arrays.asList("container-app-1", "container-app-2")),
                com.azure.core.util.Context.NONE);
    }
}
