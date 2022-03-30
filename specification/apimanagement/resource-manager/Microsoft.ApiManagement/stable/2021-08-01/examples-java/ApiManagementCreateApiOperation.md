Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.apimanagement.models.RepresentationContract;
import com.azure.resourcemanager.apimanagement.models.RequestContract;
import com.azure.resourcemanager.apimanagement.models.ResponseContract;
import java.util.Arrays;

/** Samples for ApiOperation CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateApiOperation.json
     */
    /**
     * Sample code: ApiManagementCreateApiOperation.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateApiOperation(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiOperations()
            .define("newoperations")
            .withExistingApi("rg1", "apimService1", "PetStoreTemplate2")
            .withDisplayName("createUser2")
            .withMethod("POST")
            .withUrlTemplate("/user1")
            .withTemplateParameters(Arrays.asList())
            .withDescription("This can only be done by the logged in user.")
            .withRequest(
                new RequestContract()
                    .withDescription("Created user object")
                    .withQueryParameters(Arrays.asList())
                    .withHeaders(Arrays.asList())
                    .withRepresentations(
                        Arrays
                            .asList(
                                new RepresentationContract()
                                    .withContentType("application/json")
                                    .withSchemaId("592f6c1d0af5840ca8897f0c")
                                    .withTypeName("User"))))
            .withResponses(
                Arrays
                    .asList(
                        new ResponseContract()
                            .withStatusCode(200)
                            .withDescription("successful operation")
                            .withRepresentations(
                                Arrays
                                    .asList(
                                        new RepresentationContract().withContentType("application/xml"),
                                        new RepresentationContract().withContentType("application/json")))
                            .withHeaders(Arrays.asList())))
            .create();
    }
}
```
