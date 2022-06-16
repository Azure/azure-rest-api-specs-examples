import com.azure.resourcemanager.botservice.models.BotProperties;
import com.azure.resourcemanager.botservice.models.Kind;
import com.azure.resourcemanager.botservice.models.MsaAppType;
import com.azure.resourcemanager.botservice.models.PublicNetworkAccess;
import com.azure.resourcemanager.botservice.models.Sku;
import com.azure.resourcemanager.botservice.models.SkuName;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Bots Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/CreateBot.json
     */
    /**
     * Sample code: Create Bot.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void createBot(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager
            .bots()
            .define("samplebotname")
            .withRegion("West US")
            .withExistingResourceGroup("OneResourceGroupName")
            .withTags(mapOf("tag1", "value1", "tag2", "value2"))
            .withProperties(
                new BotProperties()
                    .withDisplayName("The Name of the bot")
                    .withDescription("The description of the bot")
                    .withIconUrl("http://myicon")
                    .withEndpoint("http://mybot.coffee")
                    .withMsaAppType(MsaAppType.USER_ASSIGNED_MSI)
                    .withMsaAppId("exampleappid")
                    .withMsaAppTenantId("exampleapptenantid")
                    .withMsaAppMsiResourceId(
                        "/subscriptions/foo/resourcegroups/bar/providers/microsoft.managedidentity/userassignedidentities/sampleId")
                    .withDeveloperAppInsightKey("appinsightskey")
                    .withDeveloperAppInsightsApiKey("appinsightsapikey")
                    .withDeveloperAppInsightsApplicationId("appinsightsappid")
                    .withLuisAppIds(Arrays.asList("luisappid1", "luisappid2"))
                    .withLuisKey("luiskey")
                    .withIsCmekEnabled(true)
                    .withCmekKeyVaultUrl("https://myCmekKey")
                    .withPublicNetworkAccess(PublicNetworkAccess.ENABLED)
                    .withDisableLocalAuth(true)
                    .withSchemaTransformationVersion("1.0"))
            .withSku(new Sku().withName(SkuName.S1))
            .withKind(Kind.SDK)
            .withEtag("etag1")
            .create();
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
