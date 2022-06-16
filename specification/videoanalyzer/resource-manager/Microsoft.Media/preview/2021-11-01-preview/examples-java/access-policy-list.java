import com.azure.core.util.Context;

/** Samples for AccessPolicies List. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/access-policy-list.json
     */
    /**
     * Sample code: Lists access policy entities.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void listsAccessPolicyEntities(com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager.accessPolicies().list("testrg", "testaccount2", 2, Context.NONE);
    }
}
