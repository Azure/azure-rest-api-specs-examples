import com.azure.core.util.Context;
import com.azure.resourcemanager.videoanalyzer.models.VideoAnalyzer;
import java.util.HashMap;
import java.util.Map;

/** Samples for VideoAnalyzers Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/video-analyzer-accounts-update.json
     */
    /**
     * Sample code: Update a Video Analyzer accounts.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void updateAVideoAnalyzerAccounts(
        com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        VideoAnalyzer resource =
            manager.videoAnalyzers().getByResourceGroupWithResponse("contoso", "contosotv", Context.NONE).getValue();
        resource.update().withTags(mapOf("key1", "value3")).apply();
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
