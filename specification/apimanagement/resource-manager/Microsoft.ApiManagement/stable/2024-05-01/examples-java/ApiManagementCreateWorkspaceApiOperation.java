
import com.azure.resourcemanager.apimanagement.fluent.models.OperationContractInner;
import com.azure.resourcemanager.apimanagement.models.RepresentationContract;
import com.azure.resourcemanager.apimanagement.models.RequestContract;
import com.azure.resourcemanager.apimanagement.models.ResponseContract;
import java.util.Arrays;

/**
 * Samples for WorkspaceApiOperation CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateWorkspaceApiOperation.json
     */
    /**
     * Sample code: ApiManagementCreateWorkspaceApiOperation.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementCreateWorkspaceApiOperation(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApiOperations().createOrUpdateWithResponse("rg1", "apimService1", "wks1", "PetStoreTemplate2",
            "newoperations",
            new OperationContractInner().withDisplayName("createUser2").withMethod("POST").withUrlTemplate("/user1")
                .withTemplateParameters(Arrays.asList()).withDescription("This can only be done by the logged in user.")
                .withRequest(new RequestContract().withDescription("Created user object")
                    .withQueryParameters(Arrays.asList()).withHeaders(Arrays.asList())
                    .withRepresentations(Arrays.asList(new RepresentationContract().withContentType("application/json")
                        .withSchemaId("592f6c1d0af5840ca8897f0c").withTypeName("User"))))
                .withResponses(Arrays.asList(new ResponseContract().withStatusCode(200)
                    .withDescription("successful operation")
                    .withRepresentations(Arrays.asList(new RepresentationContract().withContentType("application/xml"),
                        new RepresentationContract().withContentType("application/json")))
                    .withHeaders(Arrays.asList()))),
            null, com.azure.core.util.Context.NONE);
    }
}
