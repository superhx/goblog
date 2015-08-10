var gulp = require('gulp');
var rename = require('gulp-rename');
var uglify = require('gulp-uglify');
var shell = require('gulp-shell');
var minifyCSS = require('gulp-minify-css');
var watch = require('gulp-watch');

var DATA_JS_SRC = 'src/data/js/';
var DATA_JS_DEST = 'data/js/';
var DATA_CSS_SRC = 'src/data/css/';
var DATA_CSS_DEST = 'data/css';
var TMPL = 'template/'
var BLOG_DIR = '/Users/Lancer/blog_test';


gulp.task('js', function () {
    gulp.src(DATA_JS_SRC + '*.js')
        .pipe(uglify())
        .pipe(rename({extname: '.min.js'}))
        .pipe(gulp.dest(DATA_JS_DEST))
        .on('end', blogReloadResource);
})

gulp.task('css', function () {
    gulp.src(DATA_CSS_SRC + '*.css')
        .pipe(minifyCSS())
        .pipe(rename({extname: '.min.css'}))
        .pipe(gulp.dest(DATA_CSS_DEST))
        .on('end', blogReloadResource);
})

gulp.task('template', shell.task([
    'rm -f category.json',
    'goblog g'
], {'cwd': BLOG_DIR}))

gulp.task('init', blogReloadResource)

gulp.task('monitor', function () {
    gulp.watch(DATA_JS_SRC + '*.js', ['js'])
    gulp.watch(DATA_CSS_SRC + '*.css', ['css'])
    gulp.watch(TMPL + '*.htm', ['template'])
})

gulp.task('default', ['js', 'css', 'template', 'monitor'])

function blogReloadResource() {
    gulp.src('').pipe(shell([
        'goblog i'
    ], {'cwd': BLOG_DIR}));
}