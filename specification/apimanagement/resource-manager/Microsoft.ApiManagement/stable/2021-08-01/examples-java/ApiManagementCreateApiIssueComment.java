import java.time.OffsetDateTime;

/** Samples for ApiIssueComment CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateApiIssueComment.json
     */
    /**
     * Sample code: ApiManagementCreateApiIssueComment.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateApiIssueComment(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiIssueComments()
            .define("599e29ab193c3c0bd0b3e2fb")
            .withExistingIssue("rg1", "apimService1", "57d1f7558aa04f15146d9d8a", "57d2ef278aa04f0ad01d6cdc")
            .withText("Issue comment.")
            .withCreatedDate(OffsetDateTime.parse("2018-02-01T22:21:20.467Z"))
            .withUserId(
                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/1")
            .create();
    }
}
