import com.azure.resourcemanager.apimanagement.models.Protocol;
import java.util.Arrays;

/** Samples for Api CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateApiNewVersionUsingExistingApi.json
     */
    /**
     * Sample code: ApiManagementCreateApiNewVersionUsingExistingApi.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateApiNewVersionUsingExistingApi(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apis()
            .define("echoapiv3")
            .withExistingService("rg1", "apimService1")
            .withSourceApiId(
                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/echoPath")
            .withDisplayName("Echo API2")
            .withServiceUrl("http://echoapi.cloudapp.net/api")
            .withPath("echo2")
            .withProtocols(Arrays.asList(Protocol.HTTP, Protocol.HTTPS))
            .withDescription("Create Echo API into a new Version using Existing Version Set and Copy all Operations.")
            .withApiVersion("v4")
            .withIsCurrent(true)
            .withApiVersionSetId(
                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apiVersionSets/aa9c59e6-c0cd-4258-9356-9ca7d2f0b458")
            .withSubscriptionRequired(true)
            .create();
    }
}
